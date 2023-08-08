import contextlib
import datetime
import os
import subprocess
import sys
import tempfile
import yaml
from git import Repo
import wget

# Download and link swagger-ui files
if not os.path.isdir('.sphinx/deps/swagger-ui'):
    Repo.clone_from('https://github.com/swagger-api/swagger-ui', '.sphinx/deps/swagger-ui', depth=1)

os.makedirs('.sphinx/_static/swagger-ui/', exist_ok=True)

if not os.path.islink('.sphinx/_static/swagger-ui/swagger-ui-bundle.js'):
    os.symlink('../../deps/swagger-ui/dist/swagger-ui-bundle.js', '.sphinx/_static/swagger-ui/swagger-ui-bundle.js')
if not os.path.islink('.sphinx/_static/swagger-ui/swagger-ui-standalone-preset.js'):
    os.symlink('../../deps/swagger-ui/dist/swagger-ui-standalone-preset.js', '.sphinx/_static/swagger-ui/swagger-ui-standalone-preset.js')
if not os.path.islink('.sphinx/_static/swagger-ui/swagger-ui.css'):
    os.symlink('../../deps/swagger-ui/dist/swagger-ui.css', '.sphinx/_static/swagger-ui/swagger-ui.css')

# Project config.
project = "LXD"
author = "LXD contributors"
copyright = "2014-%s %s" % (datetime.date.today().year, author)

with open("../shared/version/flex.go") as fd:
    version = fd.read().split("\n")[-2].split()[-1].strip("\"")

# Extensions.
extensions = [
    "myst_parser",
    "sphinx_design",
    "sphinx_tabs.tabs",
    "sphinx_reredirects",
    "sphinxext.opengraph",
    "youtube-links",
    "related-links",
    "custom-rst-roles",
    "sphinxcontrib.jquery",
    "sphinx_copybutton",
    "sphinx.ext.intersphinx",
    "terminal-output",
    "config-options",
    "notfound.extension"
]

myst_enable_extensions = [
    "substitution",
    "deflist",
    "linkify"
]

myst_linkify_fuzzy_links=False
myst_heading_anchors = 7

if os.path.exists("./substitutions.yaml"):
    with open("./substitutions.yaml", "r") as fd:
        myst_substitutions = yaml.safe_load(fd.read())
if os.path.exists("./related_topics.yaml"):
    with open("./related_topics.yaml", "r") as fd:
        myst_substitutions.update(yaml.safe_load(fd.read()))

intersphinx_mapping = {
    'cloud-init': ('https://cloudinit.readthedocs.io/en/latest/', None)
}

notfound_urls_prefix = "/lxd/en/latest/"

# Setup theme.
templates_path = [".sphinx/_templates"]

html_theme = "furo"
html_show_sphinx = False
html_last_updated_fmt = ""
html_favicon = ".sphinx/_static/favicon.ico"
html_static_path = ['.sphinx/_static']
html_css_files = ['custom.css']
html_js_files = ['header-nav.js']
html_extra_path = ['.sphinx/_extra']

html_theme_options = {
    "sidebar_hide_name": True,
    "light_css_variables": {
        "font-stack": "Ubuntu, -apple-system, Segoe UI, Roboto, Oxygen, Cantarell, Fira Sans, Droid Sans, Helvetica Neue, sans-serif",
        "font-stack--monospace": "Ubuntu Mono, Consolas, Monaco, Courier, monospace",
        "color-foreground-primary": "#111",
        "color-foreground-secondary": "var(--color-foreground-primary)",
        "color-foreground-muted": "#333",
        "color-background-secondary": "#FFF",
        "color-background-hover": "#f2f2f2",
        "color-brand-primary": "#111",
        "color-brand-content": "#06C",
        "color-api-background": "#cdcdcd",
        "color-inline-code-background": "rgba(0,0,0,.03)",
        "color-sidebar-link-text": "#111",
        "color-sidebar-item-background--current": "#ebebeb",
        "color-sidebar-item-background--hover": "#f2f2f2",
        "toc-font-size": "var(--font-size--small)",
        "color-admonition-title-background--note": "var(--color-background-primary)",
        "color-admonition-title-background--tip": "var(--color-background-primary)",
        "color-admonition-title-background--important": "var(--color-background-primary)",
        "color-admonition-title-background--caution": "var(--color-background-primary)",
        "color-admonition-title--note": "#24598F",
        "color-admonition-title--tip": "#24598F",
        "color-admonition-title--important": "#C7162B",
        "color-admonition-title--caution": "#F99B11",
        "color-highlighted-background": "#EbEbEb",
        "color-link-underline": "var(--color-background-primary)",
        "color-link-underline--hover": "var(--color-background-primary)",
        "color-version-popup": "#772953",
        "color-orange": "#FBDDD2",
    },
    "dark_css_variables": {
        "color-foreground-secondary": "var(--color-foreground-primary)",
        "color-foreground-muted": "#CDCDCD",
        "color-background-secondary": "var(--color-background-primary)",
        "color-background-hover": "#666",
        "color-brand-primary": "#fff",
        "color-brand-content": "#06C",
        "color-sidebar-link-text": "#f7f7f7",
        "color-sidebar-item-background--current": "#666",
        "color-sidebar-item-background--hover": "#333",
        "color-admonition-background": "transparent",
        "color-admonition-title-background--note": "var(--color-background-primary)",
        "color-admonition-title-background--tip": "var(--color-background-primary)",
        "color-admonition-title-background--important": "var(--color-background-primary)",
        "color-admonition-title-background--caution": "var(--color-background-primary)",
        "color-admonition-title--note": "#24598F",
        "color-admonition-title--tip": "#24598F",
        "color-admonition-title--important": "#C7162B",
        "color-admonition-title--caution": "#F99B11",
        "color-highlighted-background": "#666",
        "color-link-underline": "var(--color-background-primary)",
        "color-link-underline--hover": "var(--color-background-primary)",
        "color-version-popup": "#F29879",
        "color-orange": "#E95420",
    },
}

html_context = {
    "github_url": "https://github.com/canonical/lxd",
    "github_version": "main",
    "github_folder": "/doc/",
    "github_filetype": "md",
    "discourse_prefix": {
        "lxc": "https://discuss.linuxcontainers.org/t/",
        "ubuntu": "https://discourse.ubuntu.com/t/"}
}

source_suffix = ".md"

# List of patterns, relative to source directory, that match files and
# directories to ignore when looking for source files.
# This pattern also affects html_static_path and html_extra_path.
exclude_patterns = ['html', 'README.md', '.sphinx', 'config_options_cheat_sheet.md']

# Open Graph configuration

ogp_site_url = "https://documentation.ubuntu.com/lxd/en/latest/"
ogp_site_name = "LXD documentation"
ogp_image = "https://documentation.ubuntu.com/lxd/en/latest/_static/tag.png"

# Links to ignore when checking links

linkcheck_ignore = [
    'https://127.0.0.1:8443/1.0',
    'https://web.libera.chat/#lxd'
]

# Setup redirects (https://documatt.gitlab.io/sphinx-reredirects/usage.html)
redirects = {
    "howto/instances_snapshots/index": "../instances_backup/",
    "reference/network_external/index": "../networks/",
}


if ("TOPICAL" in os.environ) and (os.environ["TOPICAL"] == "True"):
    root_doc = "index_topical"
    exclude_patterns.extend(['index.md','tutorial/index.md','howto/index.md','explanation/index.md','reference/index.md','howto/troubleshoot.md'])
    tags.add('topical')
else:
    exclude_patterns.extend(['index_topical.md','security.md','external_resources.md','reference/network_external.md'])
    tags.add('diataxis')


@contextlib.contextmanager
def pushd(new_dir):
    previous_dir = os.getcwd()
    os.chdir(new_dir)
    try:
        yield
    finally:
        os.chdir(previous_dir)


def generate_go_docs(app):
    """
        This function calls the `lxd-doc` tool to generate
        the documentation elements from an annotated Golang codebase.
    """
    try:
        subprocess.run(["go", "version"], check=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
    except subprocess.CalledProcessError:
        raise ValueError("go is not installed for lxd-doc installation.")

    os.environ['CGO_ENABLED'] = '0'

    with tempfile.TemporaryDirectory() as tempdir:
        if os.getcwd().endswith("/doc"):
            lxddocDir = "../lxd/config/generate"
            lxdBaseDir = ".."
            outputBaseDir = ""
        else:
            lxddocDir = "lxd/config/generate"
            lxdBaseDir = "."
            outputBaseDir = "./doc/"

        with pushd(lxddocDir):
            try:
                subprocess.run(["go", "build", "-o", os.path.join(tempdir, "lxd-doc")], check=True)
            except subprocess.CalledProcessError:
                raise ValueError("Building lxd-doc failed.")

        # Generate the documentation
        try:
            subprocess.run([os.path.join(tempdir, "lxd-doc"), lxdBaseDir, "-y", f"{outputBaseDir}config_options.yaml", "-t", f"{outputBaseDir}config_options.txt"], check=True)
        except subprocess.CalledProcessError:
            raise ValueError("Generating the codebase documentation failed.")

        print("Codebase documentation generated successfully")


def setup(app):
    app.connect('builder-inited', generate_go_docs)
