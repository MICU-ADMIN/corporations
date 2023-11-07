// Define a class named "Hadith" within the "api" namespace.
namespace api
{
    // This class represents hadith data.

    public class Hadith
    {
        public int Id { get; set; }

        public string? CollectionId { get; set; }

        public string? BookId { get; set; }

        public int? HadithNumber { get; set; }

        public string? Label { get; set; }

        public string? Arabic { get; set; }

        public string? EnglishTrans { get; set; }

        public string? PrimaryNarrator { get; set; }
    }
}
