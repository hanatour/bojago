# bojago
보자고!!!
## Phase 1 계획
- 최근 배포 이력 목록으로 보여주기
- 다중 계정 고려
- 인스턴스 마련해서 자동 배포 설정하기
- 새로 고침 버튼 또는 자동 갱신
- 멀티 프로세싱 적용
- 아이디어: fail 의 경우 오토스케일링 그룹의 인스턴스 목록의 콘솔 페이지를 보여주기 -> phase 2로
- 아이디어: 상세 페이지로 콘솔 링크(P1) 또는 상세 데이터를 보여주기(P2)
- 진행 방법
  - 지금 처럼
  - 돌아가며 메인, 진행은 함께
  - 진짜 프로젝트 처럼
  - 이슈 적극 활용
- 배포 방법
  - hnt-cloud에 고정인스턴스 1ea 마련
  - 도메인 등록 bojago.hanatour.com
  - github actions / aws code pipeline / aws beanstalk / ...
  - fargate 인스턴스 사용
- 테스트도 엄격하게 적용
- #1 코드 정리 해결
