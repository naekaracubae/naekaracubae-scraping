# naekaracubae-scraping
네,카라쿠배 어플리케이션의 채용 정보를 수집하는 서비스입니다.

## 사용 기술
- golang
- serverless
- AWS Lambda
- AWS SES
- AWS RDS

## 배포 파이프라인
구축 계획중입니다.

## 작동 그림
![image](https://user-images.githubusercontent.com/81010357/137237433-e76e3b57-6e08-4b7f-8d8e-21b9501e51e0.png)
- cloudwatch를 이용해 평일 정해진 시간에 scraping과 메일 전송을 수행하는 cronjob 을 걸어놓음.
- lambda에는 scraping, 메일 전송을 수행하는 business logic 을 golang으로 제작.
- scraping 시 lambda 함수가 인터넷 엑세스 권한을 얻으려면 VPC에 NAT 게이트웨이가 있어야 한다는 이슈가 있어, [함수 앞단에 EC2를 이용해 NAT 서버를 구축](https://velog.io/@msyhu/%EC%8B%9C%EB%A6%AC%EC%A6%883-AWS%EB%A5%BC-%EC%9D%B4%EC%9A%A9%ED%95%9C-%ED%81%B4%EB%9D%BC%EC%9A%B0%EB%93%9C-%EB%84%A4%EC%9D%B4%ED%8B%B0%EB%B8%8C-%EC%96%B4%ED%94%8C%EB%A6%AC%EC%BC%80%EC%9D%B4%EC%85%98-%EC%9D%B8%ED%94%84%EB%9D%BC-%EA%B5%AC%EC%B6%95-5-NAT-Gateway-NAT-Instance-%EB%A1%9C-%EB%8C%80%EC%B2%B4%ED%95%B4%EC%84%9C-%EB%B9%84%EC%9A%A9-%EC%A0%88%EC%95%BD%ED%95%98%EA%B8%B0-step-by-step). 
