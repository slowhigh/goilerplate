package entity

// ShouldBindJSON 이기 때문에 json 형식을 적는것이다.
type Person struct {
	FirstName string `json:"firstname" binding:"required"` // "required" => 필수 항목을 의미함
	LastName  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" validate:"required,gte=1,lte=120"` // "gte"는 최소값, "lte"는 최대값, "gte=1,lte=130"는 1이상 100이하의 값
	Email     string `json:"email" binding:"required,email"` // "email"은 이메일 형식 유효성
}

// lt	: Less Than				(미만)
// lte	: Less Than Equals		(이하)
// gt	: Greater Than			(초과)
// gte	: Greater Than Equals	(이상)
// eq	: Equals				(일치)
// ne	: Not Equals			(일치X)		ex) ne=5 이면 5라는 숫자만 안되고 55는 됨, 완벽하게 일치하지 않으면 됨
// 더 많은 참고 https://pkg.go.dev/github.com/go-playground/validator#hdr-Required

type Video struct {
	Title       string `json:"titie" binding:"min=2,max=10" validate:"is-cool"` // "is-cool"이라는 이름의 사용자 유효성 검사 사용 video-controller.go 파일에 New메소드 참조
	Description string `json:"description" binding:"max=20"` // "min"은 최소 길이, "max"는 최대 길이
	URL         string `json:"url" binding:"required,url"` //"url"은 url 형식 유효성
	Author      Person `json:"author" binding:"required"`
}


// binding:"required,gte=1,lte=130" 도 작동하고 validate:"required,gte=1,lte=130" 도 작동한다. 왜 그럴까? 왜 둘다 사용할까? 두 방법의 차이는 무엇일까?
// 아직 추측이지만 binding tag는 gin에서 지원하는 거고 validate tag는 validator에서 지원하는거 같다
// gin은 binding으로 받은 값을 validator에 validate로 전환해 주는것 같다.
// gin에서 validator 패키지를 사용할 것 같다는 것이 나의 추측이다.
// 하지만 위에 있는 "is-cool"같은 커스텀 유효성 검사는 validate.RegisterValidation("is-cool", validators.ValidateCoolTitle) 이렇게 validator로 등록을 해야 하기 때문에
// binding이 아닌 validate를 사용한것 같다.
// 그래서 기존에 지원하는 유효성 검사 태그는 binding으로 사용하고, 커스텀 유효성 검사 태그는 validate로 사용하는 것이
// 코드 가독성과 유지보수성 쪽에서 좋을 것 같기도 하다.
