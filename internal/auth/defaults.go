package auth

// Allow is an auth controller which allows access to all connections and topics.
type Allow struct{}

// Auth returns true if a username and password are acceptable. Allow always
// returns true.
func (a *Allow) Authenticate(user, password []byte) bool {
	return true
}

// ACL returns true if a user has access permissions to read or write on a topic.
// Allow always returns true.
func (a *Allow) ACL(user, topic string, write bool) bool {
	return true
}

// Disallow is an auth controller which disallows access to all connections and topics.
type Disallow struct{}

// Auth returns true if a username and password are acceptable. Disallow always
// returns false.
func (d *Disallow) Authenticate(user, password []byte) bool {
	return false
}

// ACL returns true if a user has access permissions to read or write on a topic.
// Disallow always returns false.
func (d *Disallow) ACL(user, topic string, write bool) bool {
	return false
}
