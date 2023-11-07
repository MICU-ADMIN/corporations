using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Logging;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace api.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class HadithController : ControllerBase
    {
        private readonly HadithDbContext _dbContext;
        private readonly ILogger<HadithController> _logger;

        public HadithController(HadithDbContext dbContext, ILogger<HadithController> logger)
        {
            _dbContext = dbContext;
            _logger = logger;
        }


        [HttpGet("{bookId}/{collectionId}")]
        public async Task<ActionResult<IEnumerable<Hadith>>> Get(string bookId, string collectionId)
        {

            if (!string.IsNullOrEmpty(bookId) && !string.IsNullOrEmpty(collectionId))
            {
                try
                {
                    var hadithList = await _dbContext.Hadith
                        .Where(h => h.CollectionId == collectionId && h.BookId == bookId)
                        .ToListAsync();

                    return Ok(hadithList);
                }
                catch (Exception ex)
                {
                    _logger.LogError(ex, "An error occurred while retrieving Hadith.");
                    return StatusCode(500, "An error occurred while processing your request.");
                }
            }
            else
            {
                _logger.LogError( "Your request was malformed try again with the correct parameters.");
                return StatusCode(500, "An error occurred while processing your request.");
            }




        }
    }
}




//The code you provided is a C# class named "HadithController" within the "api.Controllers" namespace.
//This class is designed to work as a controller in an ASP.NET Core MVC application. Here's a breakdown
//of what this code does:
//It imports the necessary namespace for ASP.NET Core MVC controllers, which is
//Microsoft.AspNetCore.Mvc.
//The class is decorated with the[ApiController] attribute, indicating that it's an API controller.
//It specifies the route for this controller using the [Route]
//attribute.In this case, the route is based on the controller name, so requests to this controller
//will be routed to[base_url]/Hadith.
//Inside the controller class, it defines an array of weather summaries.
//It creates a logger instance for this controller using the ILogger<HadithController>.
//The constructor for the controller takes an ILogger<HadithController> as a parameter to initialize
//the logger.
//It defines an HTTP GET endpoint using the [HttpGet] attribute.This endpoint is accessible through a
//GET request to the specified route.It returns an IEnumerable of Hadith, which means it will return a
//list of weather forecasts.
//Inside the Get method, it generates a list of weather forecasts using random data.It creates 5 weather
//forecast objects with random dates, temperature in degrees Celsius, and summaries, and returns them as an
//array.
//This controller is designed to handle GET requests for weather forecasts and return a list of weather
//forecasts in response.It's part of an ASP.NET Core MVC application and follows RESTful API conventions.