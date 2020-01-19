package pkg


// Connection defines an SSH connection
type Connection struct {
		ID 					int
		HostName		string
		IPAddress		string
		AuthType		string
		AuthData		map[string]string

}