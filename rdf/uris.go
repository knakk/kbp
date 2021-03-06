package rdf

// Commonly used URIs
var (
	// rdf: https://www.w3.org/1999/02/22-rdf-syntax-ns#
	RDFbag        = NamedNode{name: "http://www.w3.org/1999/02/22-rdf-syntax-ns#Bag"}
	RDFfirst      = NamedNode{name: "http://www.w3.org/1999/02/22-rdf-syntax-ns#first"}
	RDFlangString = NamedNode{name: "http://www.w3.org/1999/02/22-rdf-syntax-ns#langString"}
	RDFnil        = NamedNode{name: "http://www.w3.org/1999/02/22-rdf-syntax-ns#nil"}
	RDFrest       = NamedNode{name: "http://www.w3.org/1999/02/22-rdf-syntax-ns#rest"}
	RDFseq        = NamedNode{name: "http://www.w3.org/1999/02/22-rdf-syntax-ns#Seq"}
	RDFtype       = NamedNode{name: "http://www.w3.org/1999/02/22-rdf-syntax-ns#type"}
	RDFHTML       = NamedNode{name: "http://www.w3.org/1999/02/22-rdf-syntax-ns#HTML"}

	// rdfs: https://www.w3.org/2000/01/rdf-schema#
	RDFScomment = NamedNode{name: "https://www.w3.org/2000/01/rdf-schema#comment"}
	RDFSdomain  = NamedNode{name: "https://www.w3.org/2000/01/rdf-schema#domain"}
	RDFSlabel   = NamedNode{name: "https://www.w3.org/2000/01/rdf-schema#label"}
	RDFSrange   = NamedNode{name: "https://www.w3.org/2000/01/rdf-schema#range"}
	RDFSseeAlso = NamedNode{name: "https://www.w3.org/2000/01/rdf-schema#seeAlso"}

	// owl: https://www.w3.org/2002/07/owl#
	OWLallValuesFrom  = NamedNode{name: "https://www.w3.org/2002/07/owl#allValuesFrom"}
	OWLcardinality    = NamedNode{name: "https://www.w3.org/2002/07/owl#cardinality"}
	OWLmaxCardinality = NamedNode{name: "https://www.w3.org/2002/07/owl#maxCardinality"}
	OWLminCardinality = NamedNode{name: "https://www.w3.org/2002/07/owl#minCardinality"}
	OWLsameAs         = NamedNode{name: "https://www.w3.org/2002/07/owl#sameAs"}
	OWLsomeValuesFrom = NamedNode{name: "https://www.w3.org/2002/07/owl#someValuesFrom"}

	// xsd: http://www.w3.org/2001/XMLSchema#
	XSDboolean       = NamedNode{name: "http://www.w3.org/2001/XMLSchema#boolean"}
	XSDbyte          = NamedNode{name: "http://www.w3.org/2001/XMLSchema#byte"}
	XSDdateTimeStamp = NamedNode{name: "http://www.w3.org/2001/XMLSchema#dateTimeStamp"}
	XSDdouble        = NamedNode{name: "http://www.w3.org/2001/XMLSchema#double"}
	XSDfloat         = NamedNode{name: "http://www.w3.org/2001/XMLSchema#float"}
	XSDgYear         = NamedNode{name: "http://www.w3.org/2001/XMLSchema#gYear"}
	XSDint           = NamedNode{name: "http://www.w3.org/2001/XMLSchema#int"}
	XSDinteger       = NamedNode{name: "http://www.w3.org/2001/XMLSchema#integer"}
	XSDlong          = NamedNode{name: "http://www.w3.org/2001/XMLSchema#long"}
	XSDshort         = NamedNode{name: "http://www.w3.org/2001/XMLSchema#short"}
	XSDstring        = NamedNode{name: "http://www.w3.org/2001/XMLSchema#string"}
	XSDunsignedByte  = NamedNode{name: "http://www.w3.org/2001/XMLSchema#unsignedByte"}
	XSDunsignedInt   = NamedNode{name: "http://www.w3.org/2001/XMLSchema#unsignedInt"}
	XSDunsignedLong  = NamedNode{name: "http://www.w3.org/2001/XMLSchema#unsignedLong"}
	XSDunsignedShort = NamedNode{name: "http://www.w3.org/2001/XMLSchema#unsignedShort"}

	// sh: http://www.w3.org/ns/shacl#
	SHpath        = NamedNode{name: "http://www.w3.org/ns/shacl#path"}
	SHname        = NamedNode{name: "http://www.w3.org/ns/shacl#name"}
	SHminCount    = NamedNode{name: "http://www.w3.org/ns/shacl#minCount"}
	SHMaxCount    = NamedNode{name: "http://www.w3.org/ns/shacl#maxCount"}
	SHclass       = NamedNode{name: "http://www.w3.org/ns/shacl#class"}
	SHdatatype    = NamedNode{name: "http://www.w3.org/ns/shacl#datatype"}
	SHNodeShape   = NamedNode{name: "http://www.w3.org/ns/shacl#NodeShape"}
	SHproperty    = NamedNode{name: "http://www.w3.org/ns/shacl#property"}
	SHtargetClass = NamedNode{name: "http://www.w3.org/ns/shacl#targetClass"}
	SHIRI         = NamedNode{name: "http://www.w3.org/ns/shacl#IRI"}
)
