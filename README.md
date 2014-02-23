# SRT - SubReddit Tracking

This has a few parts to it

## Getting a list of new subreddits

Every 30 minutes or so check for new subreddits that are not
yet accounted for and add them to the list of subreddits

## Getting a count of subscribers for each subreddit

The reddit API asks that we make only 30 API calls / minute so if this app
is tracking 30 subreddits they can all be updated in one minute.

To make sure I don't pass 30 calls/minute, I'm going to have a pool of api
calls I'm allowed to make that gets filled at a rate of 1call/2000ms.
Then every API call will go through this pool.

Will have to watch out for backpressure.

## Calculating the "populatrity" metric

First round is the older version of the HN metric.

Given the age of the subreddit and the number of subscribers calculate a value.

    Score = (S-1) / (T+2)^G

    where,
    S = Subscriber count (-1 negates the creator)
    T = time since submission (in hours)
    G = Gravity, defaults to 1.8 (will have to play wth this) 
