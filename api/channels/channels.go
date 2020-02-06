package channels

func OK(c chan bool) bool{
	select{
		case ok:= <-c :{
			if ok{
				return true
			}
		return false
		}
	}
}