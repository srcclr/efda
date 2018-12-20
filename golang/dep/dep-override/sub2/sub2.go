package sub2

import "strconv"
import "fmt"
import "github.com/simeji/jid"

func Bar() {
    query1 := jid.NewQueryWithString("HelloYeah")
    fmt.Println("Query length: " + strconv.Itoa(query1.Length()));
}
