Core AWS Administration
-----------------------

Warning: Content based on the 2012 

### Thinking about your server Presence on AWS

Key things what you should pick
  1. Pick your Region - Asia, Europe & Etc
  2. Pick Availability Zone - ua-west-1c (something like that - depends on Region)
  3. Provision you server - Select the Operation system and configuration -medium, large & etc
  4. Configure your service - web, db, media streaming & etc
  5. Expand to other availability zone 
    (i) Purpose: serve the same service to different region
    (ii) Advantage : Reduce the latency & high avail too.
  6. Expand to other region - Giving the service to different regions

| Vertical scaling  | Horizondal Scaling |
| ------------- | ------------- |
| Traditional | AWS style  |
| Increasing the HW | Increase the instance |
| Increasing the Capacity | Shared Capacity |
| Easy to do | Typically requires planning |

In simple way - Instead of increasing the hw, simply increase the instances and setup loadbalancer for requests, Whenever there is less usage simply S-OUT some of the server.
