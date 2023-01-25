# auth-service

 본 프로젝트는 http, JWT 기반의 인증을 통한 로그인 기능을 수행하는   
인증 서비스이다.   

## Directory 구조

build/ : 빌드와 관련된 파일을 관리(Dockerfile)

cmd/ : main 어플리케이션 관리

- auth_service : auth_service 시작되는 곳
    
    main.go

docs/ : 기능 정의 및 Api document 관리

models/ : struct 구조체 선언 및 관리

database/ : database 요청 함수 관련 및 요구사항 정의

- rdb : RDB 기반 정의

    crud : 직접 요청하는 쿼리 로직

.gitignore : git ignore 파일