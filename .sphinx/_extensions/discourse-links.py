from docutils import nodes
from docutils.parsers.rst import Directive
import urllib.request, json

discourse_prefix="https://discuss.linuxcontainers.org/t/"

thelinks = []

class DiscourseLinks(Directive):

    required_arguments = 0
    optional_arguments = 0
    final_argument_whitespace = True
    has_content = True

    def run(self):
        allnodes = nodes.bullet_list()
        allnodes.set_class("discourse_links")

        self.assert_has_content()

        for line in self.content:
            link = line.split(":")
            if not len(link) == 2:
                raise self.error("Required format for Discourse links is \"123456: Title\".")
            else:
                thelinks.append((link[0],link[1]))
                para = nodes.paragraph()
                newnode = nodes.reference(refuri=discourse_prefix+link[0])
                newnode.append(nodes.Text(link[1]))
                para += newnode

                allnodes += nodes.list_item('',para)

        return [allnodes]

def html_page_context(app, pagename, templatename, context, doctree):
    if pagename in context["discourse_links"]:
        linklist = "<ul>";
        for post in context["discourse_links"][pagename]:
            linkurl = discourse_prefix+post
            with urllib.request.urlopen(linkurl+".json") as url:
                data = json.load(url)
                print(data['title'])
            linklist += "<li><a href=\""+linkurl+"\">"+data['title']+"</a></li>"
        linklist += "</ul>"
        context["testing"] = linklist
    context["pagename"] = pagename
#    print(context["discourse_links"])

# The registration function
def setup_my_func(app, pagename, templatename, context, doctree):
     # The template function
    def my_func(mystring):
        posts = mystring.strip().split(",")

        linklist = "<ul>";
        for post in posts:
            linkurl = discourse_prefix+post
            with urllib.request.urlopen(linkurl+".json") as url:
                data = json.load(url)
            linklist += "<li><a href=\""+linkurl+"\">"+data['title']+"</a></li>"
        linklist += "</ul>"


        return linklist
     # Add it to the page's context
    context['my_func'] = my_func

def setup(app):
    app.add_directive("discourse-links", DiscourseLinks)
    app.connect('html-page-context', html_page_context)
    app.connect("html-page-context", setup_my_func)

    return {
        'version': '0.1',
        'parallel_read_safe': True,
        'parallel_write_safe': True,
    }
