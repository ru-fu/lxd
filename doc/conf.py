# Configuration file for the Sphinx documentation builder.
#
# This file only contains a selection of the most common options. For a full
# list see the documentation:
# https://www.sphinx-doc.org/en/master/usage/configuration.html

# -- Path setup --------------------------------------------------------------

# If extensions (or modules to document with autodoc) are in another directory,
# add these directories to sys.path here. If the directory is relative to the
# documentation root, use os.path.abspath to make it absolute, like shown here.
#
# import os
import sys
# sys.path.insert(0, os.path.abspath('.'))


# -- Project information -----------------------------------------------------

project = 'LXD documentation'
copyright = '2021, Canonical Ltd.'
author = 'Canonical Ltd.'


# -- General configuration ---------------------------------------------------

# Add any Sphinx extension module names here, as strings. They can be
# extensions coming with Sphinx (named 'sphinx.ext.*') or your custom
# ones.
extensions = [
    "myst_parser",
    "sphinx_tabs.tabs",
    "sphinx_reredirects"
]


myst_enable_extensions = [
    "substitution",
    "deflist",
    "linkify"
]

sys.path.insert(0, "reuse")
from substitutions import myst_substitutions

# Add any paths that contain templates here, relative to this directory.
templates_path = ['_templates']

# List of patterns, relative to source directory, that match files and
# directories to ignore when looking for source files.
# This pattern also affects html_static_path and html_extra_path.
exclude_patterns = ['_build', 'Thumbs.db', '.DS_Store', 'env', 'reuse']


# -- Options for HTML output -------------------------------------------------

# The theme to use for HTML and HTML Help pages.  See the documentation for
# a list of builtin themes.
#
html_theme_path = ['.']
html_theme = 'vanilla'

# html_css_files = [
#     'styles.css',
# ]


# html_title = "your custom sidebar title"

# Add any paths that contain custom static files (such as style sheets) here,
# relative to this directory. They are copied after the builtin static files,
# so a file named "default.css" will overwrite the builtin "default.css".

# Uses global TOC for side nav instead of default local TOC
html_sidebars = {
    '**': [
        'globaltoc.html',
    ]
}
#
# theme_globaltoc_maxdepth = 2
# theme_globaltoc_includehidden = True


source_suffix = [".rst", ".md"]


redirects = {
    "redirect/index": "../configuration/",
}
