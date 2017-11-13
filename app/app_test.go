package app

import "testing"

func TestFizzBuzz_Fizz(t *testing.T) {

	if FizzBuzz(3) != "Fizz" {
		t.Error("Expected returned string to be 'Fizz'")
	}
}

func TestFizzBuzz_Buzz(t *testing.T) {
	if FizzBuzz(5) != "Buzz" {
		t.Error("Expected returned string to be 'Buzz'")
	}
}

func TestFizzBuzz_FizzBuzz(t *testing.T) {
	if FizzBuzz(15) != "FizzBuzz" {
		t.Error("Expected returned string to be 'FizzBuzz'")
	}
}

func TestFizzBuzz_Num(t *testing.T) {
	if FizzBuzz(2) != "2" {
		t.Error("Expected returned string to be '2'")
	}
}

func TestReadXml_XML(t *testing.T) {
	str := `<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet type="text/xsl" href="style.xsl"?>
<People>
  <!--These are all known people-->
  <Person name="Jon"/>
  <Person name="Sally"/>
</People>
`
	result := ReadXml()

	if result != str {
		t.Errorf("Expected returned string to be '%s'", str)
	}
}