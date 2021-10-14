This is a simple generator pattern example.

Key is to wrap non concurrent-safe operation somehow, to prevent routine lock on value get and change interaction with value generator to concurrent-safe.


### Warning
This code is not production-ready. In this example generator always returns an interface{}.

Interface type coersion is very situative operation. Try to avoid it with strict types.
