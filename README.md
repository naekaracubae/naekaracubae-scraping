# GobbyIsntFree

 해리포터의 'Dobby Is Free' 라는 대사를 이용해, 나는 아직 자유가 아니라는 것을 표현한 프로젝트 제목입니다.

# 서비스 소개

 해리포터에는 [예언자 일보](https://namu.wiki/w/%EC%98%88%EC%96%B8%EC%9E%90%20%EC%9D%BC%EB%B3%B4) 라는 마법사들의 일간 신문이 있습니다. 마법사들은 본 신문을 읽으며 하루를 시작하곤 합니다. 저는 여기서 아이디어를 얻어, 두 가지 주제에 대해 매일 아침 9시에 구독자분들의 메일로 새로운 소식을 전해주는 신문 서비스를 제공해 드리고자고자 합니다.

1. 미식가 일보 : 유명한 맛집 블로거, 유튜버 분들의 전날의 새로운 글, 영상 정보를 모아 전달해드립니다.
2. 개발 채용 일보 : 요즘 개발자들이 외치는 마법같은 단어, '네카라쿠베당토' 의 채용 정보를 모아 전달해드립니다.

# 기술 스택

사용한(할) 기술 스택은 다음과 같습니다.
![image](https://user-images.githubusercontent.com/81010357/118457665-f31eda80-b734-11eb-9da7-1045a86f92f0.png)


- Infra
    - EC2 : 현재 사용하고 있는 인프라로, 안에서 cronjob을 이용해 정해진 시간에 scrapping을 하고 메일을 보냅니다.
    - Lambda : 전환할 예정인 서비스입니다. 하루에 한 번만 돌아가면 되는 본 서비스의 특성상 하루종일 머신을 대여할 필요가 없습니다. 따라서 AWS의 FaaS 서비스인 Lambda를 이용해 특정 시간에만 함수를 실행시키도록 서비스를 리팩토링할 예정입니다.
    - RDS : AWS의 RDB 제공 서비스로서, 데이터를 저장합니다.
    - SES : AWS의 메일 전송 서비스로,  고정금액 없이 내가 보낸 메일 양에 따라서만 요금을 부과하고 무료 메일 전송 수도 일당 2000개로 충분하다고 생각되어 선정했습니다.
- Language & Framework
    - Golang : 보통 scraping 하면  python을 일반적으로 사용하지만, 병렬처리에 특화되어 있는 golang의 goroutine을 이용해 scraping 속도를 높이고자 했습니다.
    - Spring boot & React : 구독 신청 & 각 사이트 정보를 보여주는 페이지 제작에 사용합니다.
  
main 이상한문장
