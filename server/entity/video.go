package entity

type Person struct {
	FirstName string `json:"firstname" binding:"required"` // "required" => 필수 항목을 의미함
	LastName  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=100"` // "gte"는 최소값, "lte"는 최대값, "gte=1,lte=130"는 1이상 100이하의 값
	Email     string `json:"email" binding:"required,email"` // "email"은 이메일 형식 유효성
}

// lt	: Less Than				(미만)
// lte	: Less Than Equals		(이하)
// gt	: Greater Than			(초과)
// gte	: Greater Than Equals	(이상)
// ne	: Not Equals			(일치X)		ex) ne=5 이면 5라는 숫자만 안되고 55는 됨, 완벽하게 일치하지 않으면 됨



type Video struct {
	Title       string `json:"titie" binding:"min=2,max=10" validate:"is-cool"` // "is-cool"이라는 이름의 사용자 유효성 검사 사용 video-controller.go 파일에 New메소드 참조
	Description string `json:"description" binding:"max=20"` // "min"은 최소 길이, "max"는 최대 길이
	URL         string `json:"url" binding:"required,url"` //"url"은 url 형식 유효성
	Author      Person `json:"author" binding:"required"`
}
