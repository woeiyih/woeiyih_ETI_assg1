# woeiyih_ETI_assg1

#Introduction
Welcome to the Oobor microservices where we design car-sharing related microservices. 

#Features
In this, you can find 2 of our microservices, which includes 1. accessing and creating Passengers and 2. creating and accessing Drivers. 

#Requirements
You would need to have MySQL Community Edition installed. 
You would also need to have Visual Studio Code installed.

#Quick Start
What you need to do is to connect to the MySQL server database and run the main.go file

#Design Considerations and API
The design focuses on accurately checks if the users have already been in the system, such that it can POST to create a new user. 
This design, after consideration, uses REST API, as we want to take advantage of existing protocol. Other factors also includes: 1. We want flexibility for our microservices as data is not really tied to any resources or methods. 2. This design is also good as we want to allow the API to be accessible to different type of projects - Web, iOS, even IoT, etc. Essentially we are future-proofing these microservices.
