# SRT - SubReddit Tracking

This has a few parts to it

## Getting a list of new subreddits

Every 30 minutes or so check for new subreddits that are not
yet accounted for and add them to the pool

## Getting a count of subscribers for each subreddit

The reddit API asks that we make only 30 API calls / minute so if this app
is tracking 30 subreddits they can all be updated.

I'm going to have a pool of api calls i'm allowed to make that gets
filled at a rate of 1call/500ms. Then every API call will have to
grab from this pool.

## Calculating the "populatrity" metric

Given the age of the subreddit and the number of subscribers calculate a value.

    Score = (S-1) / (T+2)^G

    where,
    S = Subscriber count (-1 negates the creator)
    T = time since submission (in hours)
    G = Gravity, defaults to 1.8 (will have to play wth this) 
