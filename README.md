# Task_Home24

## Objective

An attempt to make a basic web scrapper using Golang with following functionalities

- HTML Version
- Page Title
- Headings count by level
- Amount of internal and external links
- Amount of inaccessible links
- If a page contains a login form

## How to Run

Assuming you already have git and go installed.

1. Run the following command to clone the latest code from git

   ```
   git clone https://github.com/SheryarButt/Task_Home24.git
   ```
2. Install dependancies using the following command

   ```
   cd Task_Home24
   go install
   ```
3. Run the following command to start webserver.

   ```
   go run main.go
   ```
4. Open Chrome or any browser of choice and go to the following address.

   ```
   http://localhost:8080/
   ```

   If you see a Hello World dialogue, It means the server is running successfully 🥳
5. To get counts of headings, links, HTML version and title of the page, use the following endpoint:

   ```
   http://localhost:8080/getCounts?url=LINK_TO_A_WEBPAGE
   I.E. 
   http://localhost:8080/getCounts?url=http://www.google.com
   ```
6. To get details of each headings, links, HTML version and title of the page, use the following endpoint:

   ```
   http://localhost:8080/getDetails?url=LINK_TO_A_WEBPAGE
   I.E. 
   http://localhost:8080/getDetails?url=http://www.google.com
   ```
7. A little recommendation, use the following extension for better viewability on chrome.

   ```
   https://github.com/tulios/json-viewer
   ```

## Pages used for testing

 Use the following links (if required) to verify integrity of the application:

- https://www.w3.org/People/mimasa/test/xhtml/media-types/test8.html
- https://www.w3.org/People/mimasa/test/xhtml/media-types/test7.html
- https://www.w3.org/People/mimasa/test/xhtml/media-types/test6.html
- https://www.w3.org/People/mimasa/test/xhtml/media-types/test4.html

## Hardware/Software Setup 

 - Dell Latitude 7470 
 - Intel(R) Core(TM) i5-6300U
 - 16.0 GB RAM
 - Windows 10 Pro v21H2, 64 Bit
 - Go v13.3
