import http from 'k6/http';
import { check } from 'k6';
import { sleep } from 'k6';
import accrual from 'k6/x/accrual';

export const options = {
  vus: 10000,
  duration: '15s',
};

export default function () {
  let jsonDate = accrual.Generator()
  console.log(`jsonData: ${jsonDate}`);

/*     const res = http.get('http://localhost:8081/api/orders/6642691858');
  check(res, {
    'is status 200': (r) => r.status === 200,
  }); */
  sleep(1);
}
