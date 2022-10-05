import unittest
from app import app
from unittest import TestCase

baseUrl = "/api/v1/analysis"

class TestApi(TestCase):

    def testGetAllData(self):
        tester = app.test_client(self)
        resp = tester.get(f'{baseUrl}/')
        self.assertEqual(resp.status_code,200)

    def testCreateInfromation(self):
        jsonData = {"category": "Bermain","date": "2022-10-10","duration": 29000000,"id": 18,"tags": "Belajar,Coding","total_duration": 30000000,"user_id": 1}
        tester = app.test_client(self)
        resp = tester.post(f'{baseUrl}/',json=jsonData, content_type='application/json')
        self.assertAlmostEqual(resp.status_code,201)
        self.assertAlmostEqual(resp.data.decode('ASCII'),'{"data":{"message":"success add payload"},"status":201}\n')
        
    def testGetAllData(self):
        tester = app.test_client(self)
        resp = tester.get(f'{baseUrl}/1')
        self.assertEqual(resp.status_code,411)
        self.assertAlmostEqual(resp.data.decode('ASCII'),'{"data":"411 Length Required: data must be more than 10 rows","status":411}\n')
    

if __name__ == '__main__':
    unittest.main()