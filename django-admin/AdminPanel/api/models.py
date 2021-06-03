from django.db import models
from ckeditor.fields import RichTextField

# Create your models here.
class Article(models.Model):
    title = models.CharField('Title', max_length=255, unique=True, null=False, default="Title")
    createdat = models.DateTimeField('Created at', null=False, auto_now_add=True)
    tags = [
        ('news', 'news'),
        ('analytics', 'analytics'),
        ('breaking', 'breaking'),
    ]
    tag = models.CharField('Tag', max_length=15, null=False, default="news", choices=tags)
    image_url = models.CharField('Image url', max_length=512, null=False, default="/beyondf1.jpg")
    text = RichTextField('Text', null=False, default="Text")
    source = models.CharField("Source", max_length=512, null=False, default="beyondf1.com")

    def __str__(self) -> str:
        return f"[{self.tag}] {self.title[:20]}..."