# gitrunner
gitrunner pulls all git PR requests of last week and provides summary data for the same

In config file :
1. We can add email address of sender and reciever. 
2. Change prevdays to any desired number of days to get data from github. For example currently prevdays is set to 7 (week). Hence it considers PR data upto 1 week.
3. The url of repository can also be changed.

While Running the docker container please consider that,
The application prints the summary data to console rather than sending email.


Sample output is as follows,


To: admin_devops@gmail.com
From: sender@gmail.com
Subject: Summary Report of last weeks github PRs for repo  home-assistant/developers.home-assistant
 The summary table is as follows
 
--------------------------------------
|   State of PR    |       Count      |
--------------------------------------
| closed    |       12         |

| total    |       16         |

| merged    |       11         |

| open    |       4         |

------------------------------------------------------------------------

