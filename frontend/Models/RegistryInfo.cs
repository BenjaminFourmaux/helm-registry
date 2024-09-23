using System.Text.Json.Serialization;

namespace frontend.Models
{
    public class RegistryInfo
    {
        public string Name { get; set; } = string.Empty;
        public string Description { get; set; } = string.Empty;
        public string Maintainer { get; set; } = string.Empty;
        [JsonPropertyName("maintainer_url")]
        public string MaintainerUrl { get; set; } = string.Empty;
        public string[] Labels { get; set; } = new string[0];
        [JsonPropertyName("number_of_repos")]
        public int NumberOfRepos { get; set; } = 0;

        public TagItem[] ToTagItem()
        {
            var listTags = new List<TagItem>();
            if (this.Labels.Length != 0)
            {
                foreach (string label in this.Labels)
                {
                    listTags.Add(new TagItem(label, "secondary"));
                }
            }
            return listTags.ToArray();
        }
    }
}
