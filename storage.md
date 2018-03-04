# Storage

How is stuff layed out on disk.

* Slotted page structure
    * Pages
    * Header contains a pointer to free space.
    * Header contains pointers to records, and length of records
    * Row ids are basically record pointers. They are a basically a pointer to
      the info in the page header, which is the offset to the actual record.
    * Deleting/Updating a record means changing all the record positions, and
      updating the header
    * Pages are linked together to help a table scan.
    * rowid = <page id>-<offset id>
    * null bitmap can be used to identify whether or not a field is null.
    * store a length value in the fixed size portion of a record, all variable
      length fields go at the end
    * types of page
        * free list page
        * data page
        * payload overflow page?
        * https://www.sqlite.org/fileformat.html
* transaction log involves writing out an unmodified version of the page to the
  log. If the transaction doesn't complete, the transaction log is used to
  restore the unmodified page.

