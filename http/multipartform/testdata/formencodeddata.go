package testdata

const (
	// FormDataWithArrays for testing .... arrays
	FormDataWithArrays = `--xYzZY
Content-Disposition: form-data; name="headers"

MIME-Version: 1.0
Received: by 0.0.0.0 with HTTP; Wed, 10 Aug 2016 18:10:13 -0700 (PDT)
From: Example User <test@example.com>
Date: Wed, 10 Aug 2016 18:10:13 -0700
Subject: Inbound Parse Test Data
To: inbound@inbound.example.com
Content-Type: multipart/alternative; boundary=001a113df448cad2d00539c16e89

--xYzZY
Content-Disposition: form-data; name="my_array[]"

I am a dimunitive dispenser of flavor infused water
--xYzZY
Content-Disposition: form-data; name="my_array[]"

short and stout
--xYzZY--
`
	// FormDataWithStructs for testing multipart encoded form with structs
	FormDataWithStructs = `--xYzZY
Content-Disposition: form-data; name="ptr_test_struct"

{ "Foo": "qux", "Bar": 1, "Baz": ["elm1", "elm2"] }
--xYzZY
Content-Disposition: form-data; name="test_struct"

{ "Foo": "quux", "Bar": 2, "Baz": ["1mle", "2mle"]}
--xYzZY--
`
	// FormDataErryThang for testing multipart encoded form with
	// everything we support
	FormDataErryThang = `--xYzZY
Content-Disposition: form-data; name="ptr_test_struct"

{ "Foo": "qux", "Bar": 1, "Baz": ["elm1", "elm2"] }
--xYzZY
Content-Disposition: form-data; name="array[]"

element1
--xYzZY
Content-Disposition: form-data; name="array[]"

element2
--xYzZY
Content-Disposition: form-data; name="bool"

true
--xYzZY
Content-Disposition: form-data; name="int"

12345678
--xYzZY
Content-Disposition: form-data; name="int_eight"

-1
--xYzZY
Content-Disposition: form-data; name="int_sixteen"

-512
--xYzZY
Content-Disposition: form-data; name="int_thirty_two"

-1024
--xYzZY
Content-Disposition: form-data; name="int_sixty_four"

-2048
--xYzZY
Content-Disposition: form-data; name="you_int"

1
--xYzZY
Content-Disposition: form-data; name="you_int_eight"

2
--xYzZY
Content-Disposition: form-data; name="you_int_sixteen"

4
--xYzZY
Content-Disposition: form-data; name="you_int_thirsty_two"

16
--xYzZY
Content-Disposition: form-data; name="you_int_sixty_four"

32
--xYzZY
Content-Disposition: form-data; name="root_beer_float_thirsty_two"

1.1
--xYzZY
Content-Disposition: form-data; name="root_beer_float_sixty_four"

1.2
--xYzZY
Content-Disposition: form-data; name="yarn"

spam
--xYzZY
Content-Disposition: form-data; name="map_worky"

{"um": "hello", "green": "eggs and spam"}
--xYzZY
Content-Disposition: form-data; name="map_no_worky"

{"um": "hello", "green": "eggs and spam"}
--xYzZY--`
)
