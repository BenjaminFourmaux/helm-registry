namespace frontend.Models
{
    public class TagItem
    {
        public string label { get; set; } = string.Empty;
        public string color { get; set; } = string.Empty;
        public bool? canRemoved { get; set; } = false;

        public TagItem(string label, string color)
        {
            this.label = label;
            this.color = color;
        }
    }
}
