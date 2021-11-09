---
substitutions:
  key4: "Override for **key4**"
---

# Testing Markdown in Sphinx

```{toctree}
:hidden:

test2
```

## Linking

* Link to another page: [](preseed.md), [](preseed), {doc}`preseed`

  *works with or without `.md`, without specifying a link text*

* Link to a section on another page: [doesn't work](../preseed.md#default-profile), [works](../preseed/#default-profile)

  *possible, but pretty tricky, not automatic link text*

* Link to a section with ID: {ref}`some_ID`, {ref}`just-for-testing`, {ref}`can override link text <some_ID>`

  *works, but uses a different syntax than Markdown or RST*

## Lists

Simple lists work just fine:

- A simple bullet.
- A more complex bullet.

  With several paragraphs.
- And an even more complex one.

       with a code block

Using "#." does not work:

1. A numbered list.
2. Number 2.
#. Next one.

But we can use "1." for all:

1. A numbered list.
1. Number 2.
1. Next one.


Alternatively, it can be done with `{eval-rst}`:

```{eval-rst}
1. A numbered list.
2. Number 2.
#. Next one.
```

Numbered lists with alternative numbering do not work:

a. Item
b. Item
c. Item

i. Item
ii. Item
iii. Item

It works with `{eval-rst}`, but is not supported by the theme:

```{eval-rst}
a. Item
#. Item
#. Item
```

```{eval-rst}
i. Item
#. Item
#. Item
```

### Nested lists

- Item

  - Nested
  - Nested
- Item
  - Nested
  - Nested
- Item
  * Nested
  * Nested
- Item
  - Nested 1
    - Nested 2

1. Item
   - Nested
2. Item
   1. Nested
   2. Nested
3. Item
   - Nested
     1. Nested
     2. Nested
   - Nested
     3. Nested
     4. Nested

### Definition lists

Term 1
: Definition

Term 2
: Definition

## Notes etc.

!!! note
    Markdown notes don't work.

```{note}
Standard note.
```

```{note}
Standard note.

But with more content.

    a code line

```

````{note}
A code block within a note.

  ```
  slkdflkdj
  ```
````

:::{note}
And here's a note with a colon fence!
:::

:::{note}
And here's a note with a colon fence!

More content.

    a code line

```
And a code block.
```
:::

:::{caution}
More note types.
:::

:::{tip}
More note types.
:::

:::{admonition} Important
More note types.
:::

## Reuse

### Substitutions

Defined in file `reuse/substitutions.py`, which is pulled into the conf file.

Inline: {{ key1 }}

Block level:

{{ key2 }}

{{ key3 }}

| col1     | col2     |
| -------- | -------- |
| {{key2}} | {{key1}} |

Override a key for one file: {{key4}}

### Include files

Include a markdown file:

% Include content from file [reuse/include1.md](reuse/include1.md)

---

```{include} reuse/include1.md
```
---

Having a header in this file raises a warning though - can we suppress warnings for include files?


Include an RST file:

---
```{eval-rst}
.. include:: reuse/include2.rst
```
---

## Tabs

````{tabs}

 ```{group-tab} Unordered list

    * Sed sagittis eleifend rutrum
    * Donec vitae suscipit est
    * Nulla tempor lobortis orci
 ```
 ```{group-tab} Ordered list

1. Sed sagittis eleifend rutrum
2. Donec vitae suscipit est
3. Nulla tempor lobortis orci
 ```
````


(some_ID)=
## Details

No support in RST, but can be done by inserting HTML ...

<details>
<summary><a>Details</a></summary>

* One
* Two
* Three

```
some code
```

</details>

## Glossary

Here's a term: {term}`environment`

## Tables

Standard markdown tables:

| a    | b    |
| :--- | ---: |
| c    | d    |

RST list tables:

```{list-table} Frozen Delights!
   :widths: 15 10 30
   :header-rows: 1

* - Treat
  - Quantity
  - Description
* - Albatross
  - 2.99
  - On a **stick**!
* - Crunchy Frog
  - [](preseed)
  - If we took the bones out, it wouldn't be
    crunchy, now would it?

```
