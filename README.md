# Log Service

## Terminology

1. Record: The data stored in the log
1. Store: The file we store records in
1. Index: The file we store index entries in
1. Segment: the abstraction that ties the store and index together
1. Log: The abstraction that ties all the segments together