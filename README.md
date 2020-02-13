This is a quick demo used to figure out some apparently odd behaviour with Bleve.

The answer ended up being https://github.com/blevesearch/bleve/issues/417 but I thought I might as well publish it for posterity and maybe to serve as a simple usage example for Bleve.

TLDR:
* To sort by a text field it needs a keyword mapping for a stable sort order (which might mean duplicating the field)
* When using a struct generated from a protobuf, the fields are identified using the original names from the protobuf and not the CamelCased version the protobuf compiler generates.
