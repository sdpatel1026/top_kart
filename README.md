# top_kart
A global E-commerce website named Topkart uses two kinds of APIs to handle lightning deals. Lightening deals
are products that are available at a discounted price on the website for a brief amount of time.The expiry 
time of a lightning deal is not more than 12 hours. These will be refreshed at 00:00 UTC daily. 
The functionality of these APIs:
1. Operations API - Only a single instance of this API is present. Which is used by the internal team.
    * Create and update lightning deals
    * Approve orders

2. Customer API - Multiple APIs are present across the globe. Using this API customers can:
    * Access available unexpired deals
    * Place orders 
    * Check the status of their order


***A lightning deal contains the following data points:***  
    * Product Name  
    * Actual & Final price  
    * Total & Available units  
    * Expiry Time  

***Considerations:***    
    * Users should not be able to place an order for a deal that is expired.  
    * Orders that are created before the expiry time of a deal are waitlisted until the internal team approves.  
    * If all the units of a lightning deal are sold, pending waitlisted applications should be auto-rejected.  
    * The customer API receives peak load at 00:00 UTC. High availability of API is expected.  



