# 기능 정의

1. 회원 가입 기능   
    1.1 입력 정보 :   
        아이디 (또는 이메일) : 길이 제한, 중복 제한(아이디만 바로 체크)   
        패스워드 : 길이 제한, 형식 제한   
        패스워드 확인 : 비교   
        이름 : 길이 제한   
        생년월일 : 형식 제한   
        성별 : 형식 제한   
        패스워드 찾기 질문 : 길이 제한   
        패스워드 찾기 답변 : 길이 제한   

    1.2 필요 API :   
        회원 가입 - 성공   
                    실패   
        아이디 중복 - 아이디 조회   
            있으면 ok   
            없으면 nok   

2. 로그인 기능   
    2.1 입력 정보 :   
        아이디   
        패스워드   

    2.2 필요 API :   
        로그인 - 아이디 조회,   
                없으면 nok   
                있으면 ok, password 비교   
                    안 맞으면 nok   
                    맞으면 ok, 로그인 성공   
                        access_token 및 refresh_token 생성   

3. 아이디 찾기 기능   
    3.1 입력 정보 :   
        아이디   

    3.2 필요 API :   
        아이디 중복 - 아이디 조회   
                있으면 ok   
                없으면 nok   

4. 패스워드 찾기 기능   
    4.1 입력 정보 :   
        아이디   
        패스워드 찾기 질문   
        패스워드 찾기 답변   

    4.2 필요 API :   
        패스워드 찾기 - 아이디 조회   
                없으면 nok   
                있으면 ok, 패스워드 찾기 질문, 답변 비교   
                    안 맞으면 nok   
                    맞으면 ok, 패스워드 찾기 성공   
                        access_token 및 refresh_token 생성   
        아이디 중복 - 아이디 조회   
            있으면 ok   
            없으면 nok   

5. 패스워드 변경 기능   
    5.1 입력 정보 :   
        패스워드   
        패스워드 확인   

    5.2  필요 API :   
        패스워드 유효성 체크   
        access_token - 아이디 조회   
                실패 nok   
                성공 ok - 업데이트 성공   

6. 정보 조회 기능   
    6.1 입력 정보 :
        access_token, refresh_token   

    6.2 필요 API :   
        정보 조회
        access_token - 아이디 조회   
            패스워드   
            이름   
            생년월일   
            성별   
            패스워드 찾기 질문   
            패스워드 찾기 답변   

7. 정보 업데이트 기능   
    7.1 입력 정보 :   
        access_token, refresh_token   

    7.2 필요 API :   
        정보 업데이트   
        access_token - 아이디 조회      
            이름
            패스워드   
            패스워드 확인     
            생년월일   
            성별   
            패스워드 찾기 질문   
            패스워드 찾기 답변   

8. 탈퇴 기능   
    8.1 입력 정보 :   
        access_token, refresh_token   

    8.2 필요 API :   
        탈퇴   
        access_token - 아이디 조회   
            정보 삭제   
            access_token, refresh_token 삭제   

9. 로그아웃 기능   
    9.1 입력 정보 :   
        access_token, refresh_token   
    
    9.2 필요 API :    
        로그아웃 :   
        access_token - 아이디 조회   
        access_token, refresh_token 삭제   

10. Token 생성   

11. Token 체크   

12. Token 재갱신   

13. Token 만료   

14. Token 제거   

14. 확장, 보완 - 사용자 인증   
    이메일 인증, 휴대폰 인증 고려   

---