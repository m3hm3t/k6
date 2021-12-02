import http from 'k6/http';
import sleep from 'k6';

const bigFile = open('./5MB_testdata.bin', 'br');
const binFile = open('./5MB_testdata.bin', 'b');

export default () => {
  const data = {
    field: 'This is a standard form field',
    file: http.file(bigFile, '5MB_testdata.bin'),
  }

  const res = http.post('https://example.com/upload', data);
  sleep(5);
  console.log(res);
};
