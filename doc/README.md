# LXD documentation

The LXD documentation is available at: <https://documentation.ubuntu.com/lxd/en/latest/>

GitHub provides a basic rendering of the documentation as well, but important features like includes and clickable links are missing. Therefore, we recommend reading the [published documentation](https://documentation.ubuntu.com/lxd/en/latest/).

## How it works

<!-- Include start docs -->

### Documentation framework

LXD's documentation is built with [Sphinx](https://www.sphinx-doc.org/en/master/index.html) and hosted on [Read the Docs](https://about.readthedocs.com/). <!-- wokeignore:rule=master -->

It is written in [Markdown](https://commonmark.org/) with [MyST](https://myst-parser.readthedocs.io/) extensions.
For syntax help and guidelines, see the [MyST style guide](https://canonical-documentation-with-sphinx-and-readthedocscom.readthedocs-hosted.com/style-guide-myst/). <!-- and the [documentation cheat sheet](https://documentation.ubuntu.com/lxd/en/latest/doc-cheat-sheet-myst/) ([source](https://raw.githubusercontent.com/canonical/lxd/main/doc/doc-cheat-sheet-myst.md)). -->

For structuring, the documentation uses the [Diátaxis](https://diataxis.fr/) approach.

### Build the documentation

To build the documentation, run `make doc` from the root directory of the repository.
This command installs the required tools and renders the output to the `doc/html/` directory.
To update the documentation for changed files only (without re-installing the tools), run `make doc-incremental`.

Before opening a pull request, make sure that the documentation builds without any warnings (warnings are treated as errors).
To preview the documentation locally, run `make doc-serve` and go to [`http://localhost:8001`](http://localhost:8001) to view the rendered documentation.
