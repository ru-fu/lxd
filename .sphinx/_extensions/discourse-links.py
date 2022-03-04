from docutils import nodes
from docutils.parsers.rst import Directive

discourse_prefix="https://discuss.linuxcontainers.org/t/"

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
                para = nodes.paragraph()
                newnode = nodes.reference(refuri=discourse_prefix+link[0])
                newnode.append(nodes.Text(link[1]))
                para += newnode

                allnodes += nodes.list_item('',para)

        return [allnodes]


def setup(app):
    app.add_directive("discourse-links", DiscourseLinks)

    return {
        'version': '0.1',
        'parallel_read_safe': True,
        'parallel_write_safe': True,
    }
