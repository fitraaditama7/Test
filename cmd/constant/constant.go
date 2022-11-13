package constant

import "time"

const ContextTimeout = 10 * time.Second
const CreditCardSeparator = "***"
const HeaderKey = "key"

var QueryOrderBy = []string{"name", "email"}
var QuerySortBy = []string{"asc", "desc"}
