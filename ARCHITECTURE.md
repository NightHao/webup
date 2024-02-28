# Architecture

This document describes the high-level architecture of webup.
Since webup is not available for public release yet, the document describes the
entire architecture of Taiwan AI educational platform.
This document should be of use for isolation of components for other usage.


## Components

### Frontend

The frontend lives partially in this codebase, while the entire frontend is
stored and hosted on NAS.
To help future cleanup, only files that are edited after the first revamp are
VCS-tracked.
The frontend is not compiled using any framework.

Apart from layout and visual, the frontend interfaces with two other components
directly, namely webup (The Backend), and algolia.
In the context of frontend, webup provides the functionality of dynamic content
serving. The course searching functionality is powered by algolia*.

*Note: the management account for algolia service is inaccessible at the time of
writing.

### The Backend

The primary feature of webup is to manifest a Content Management System (CMS)
leveraging the features provided by Google Drive and Google Docs, such that the
content can be maintained and managed almost entirely via Google web interface,
and a web front can be developed to render the content by interpreting and
navigating resources as returned by webup API.
For more details refer to the **code map** section.

### Algolia

The clickable dots on the interactive course map redirect user to a list of
courses for the relevant topic.
It is implemented by controlling the search parameters to algolia.
The management account is inaccessible and has never been, thus there is not
much to be documented, honestly.
This block exists merely to emphasize the current complexity of the entire
system.

## Code Map

### General Note

While the code currently lives in `internal`, all packages within are developed
with generic usage in mind, and is possible to be moved into `pkg` once the APIs
are stabilized.

### G Suites

The three packages prefixed with `g` are all packages dealing with Google APIs.
Ideally a `g` package should include and abstract functions that will be used
by different subpackages, this does not currently hold true.

### Google Docs

Webup interfaces with user via content editable using Google Docs Editors.
Thus, functionalities to retrieve, parse and process those documents are
packaged according to their subtypes, i.e. Google Docs (`gdoc`) and Google
Sheets (`gsheet`).

A convention is currently being practiced:
 
- `elements.go`: Define types returned by Google APIs
- `net.go`: HTTP related functions

### CMS

Webup aims to provide a library of functions for you to design your own CMS,
without imposing too many constraints.
An architecture invariant is that no Google-identifying-ID will be presented
to the frontend, this is to ensure that permission misconfiguration of the drive
folders and documents will not pose security threat.
There are four supported unit of contents: `DOC`, their corresponding accessor
type `DOCLIST`, `LINK` and `MENU`.

A `DOC` can be (partially) rendered into HTML, functioning as an editable
webpage.
Following the invariant, a `DOC` can only be accessed from the API using their
index into their parent `DOCLIST`.

A `DOCLIST` is a drive folder with some Gdoc documents.
Similar to a `DOC`, a `DOCLIST` can onlg be accessed using their index into
their parent `MENU`.

A `MENU` is a spreadsheet that defines a multi-level nestable dropdown menus.

The simplest implementation of an CMS can be achieved by defining a root `MENU`.

