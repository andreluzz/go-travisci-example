package xml

import "testing"

func TestParser_Success(t *testing.T) {

	str := `<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet type="text/xsl" href="style.xsl"?>
<People>
  <!--These are all known people-->
  <Person name="Jon"/>
  <Person name="Sally"/>
</People>
`
	result, err := Parser()

	if err != nil {
		t.Error("Error parsing xml")
	}

	if result != str {
		t.Errorf("Expected returned string to be '%s'", str)
	}
}
