from django.contrib import admin
from core.models import Video, Tag, VideoMedia

#class VideoMediaInline(admin.StackedInline):
#  model = VideoMedia
#  verbose_name = 'Mídia'
#  max_num = 1
#  min_num = 1
#  can_delete = False

class VideoAdmin(admin.ModelAdmin):
  #inlines = [VideoMediaInline]
  pass

# Register your models here.
admin.site.register(Video)
admin.site.register(Tag)
admin.site.register(VideoMedia)