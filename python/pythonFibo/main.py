import webapp2
import time
import timeit
import datetime
from google.appengine.ext import ndb

def Fibonacci(n) :
    if (n == 1 or n == 2):
        FiboResult = 1
    else :
        FiboResult = Fibonacci(n - 1) + Fibonacci(n - 2)
    return FiboResult

def FiboTime ( fiboIndex ):
    startTime = timeit.default_timer()
    Fibo = Fibonacci( fiboIndex )
    endTime = timeit.default_timer()
    totalTime = (endTime - startTime)*1000
    return totalTime

class Data(ndb.Model):
    FiboIndex = ndb.IntegerProperty()
    TimeStamp = ndb.DateTimeProperty()
    Results = ndb.FloatProperty()

def dataStore(totalTime, fiboIndex):
    data = Data(
        FiboIndex=fiboIndex,
        TimeStamp=datetime.datetime.utcnow(),
        Results=totalTime)
    data.put()
    return data

class MainPage(webapp2.RequestHandler):
    def get(self):
        fiboIndex = 40
        for flag in range(0, 30):
            totalTime = FiboTime(fiboIndex)
            data = dataStore(totalTime,fiboIndex)
            time.sleep(20)
        self.response.headers['Content-Type'] = 'text/plain'
        self.response.write('Hello, World!')


app = webapp2.WSGIApplication([
    ('/', MainPage),
], debug=True)
