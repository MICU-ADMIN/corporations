using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Design;
using Microsoft.Extensions.Configuration;
using System.IO;

public class HadithDbContextFactory : IDesignTimeDbContextFactory<HadithDbContext>
{
    public HadithDbContext CreateDbContext(string[] args)
    {
        IConfigurationRoot configuration = new ConfigurationBuilder()
            .SetBasePath(Directory.GetCurrentDirectory())
            .AddJsonFile("appsettings.json")
            .Build();

        var optionsBuilder = new DbContextOptionsBuilder<HadithDbContext>();
         optionsBuilder.UseMySql(configuration.GetConnectionString("HadithDatabase"));

        return new HadithDbContext(optionsBuilder.Options);
    }
}