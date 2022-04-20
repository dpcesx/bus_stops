Example schedule for

  Stop 1:
    Route 1 12:00, 12:15, 12:30, 12:45 ...
    Route 2 12:02, 12:17, 12:32, 12:47 ...
    Route 3 12:04, 12:19, 12:34, 12:49 ...
    etc..

  Stop 2
    Route 1 12:02, 12:17, 12:32, 12:47 ...
    Route 2 12:04, 12:19, 12:34, 12:49 ...
    Route 3 12:06, 12:21, 12:36, 12:51 ...
    etc..

The API should return to the consumer the
  next 2 arrival times
    per route
      for the requested stop.

Example
Get/Retrieve for Stop ID 1

Would return 2 arrival times each for route 1, 2, and 3


It should output the updated prediction times every minute until stopped.

Example output when running the app at 3:01PM:

Stop 1:
  Route 1 in 14 mins and 29 mins
  Route 2 in 1 mins and 16 mins
  Route 3 in 3 mins and 18 mins

Stop 2:
  Route 1 in 1 mins and 16 mins
  Route 2 in 3 mins and 18 mins
  Route 3 in 5 mins and 20 mins
