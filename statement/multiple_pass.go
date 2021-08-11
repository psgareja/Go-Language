const (
	usage="Usage:[username][password]"
	errUser="Access Denied %q.\n"
	errPwd="Invalid %q\n"
	accessOK="Access %q\n"
)
user,user2="jack","inace"
pass,pass2="1888","1222"
func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	if u != user && u!=user2 {
		fmt.Printf(errUser, u)
	} else if p != pass && p!=pass2 {
		fmt.Printf(errPwd, u)
	} else {
		fmt.Printf(accessOK, u)
	}
}
