package main

type Stack struct {
  source []interface{}
}

func (s *Stack) push(e interface{}) {
  s.source = append(s.source, e)
}

func (s *Stack) pop() interface{} {
  var x interface{} 
  x, s.source = s.source[len(s.source)-1], s.source[:len(s.source)-1]
  return x
}
