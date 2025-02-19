import http from 'k6/http';
import { check } from 'k6';

export let options = {
    stages: [
        { duration: '1m', target: 100000 }, // ramp up to 100k requests
        { duration: '5m', target: 100000 }, // stay at 100k requests
        { duration: '1m', target: 0 }, // ramp down
    ],
};

export default function () {
    let res = http.get('http://localhost:8080/api/surgeForecast');
    check(res, {
        'status is 200': (r) => r.status === 200,
    });
} 