using api;
using Microsoft.EntityFrameworkCore;

public class HadithDbContext : DbContext
{
    public HadithDbContext(DbContextOptions<HadithDbContext> options) : base(options)
    {
    }

    public DbSet<Hadith> Hadith { get; set; }
}