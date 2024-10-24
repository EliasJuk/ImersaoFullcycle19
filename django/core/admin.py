from django.contrib import admin
from core.models import Video, Tag, VideoMedia

# Register your models here.
admin.site.register(Video)
admin.site.register(Tag)
admin.site.register(VideoMedia)