from django.contrib import admin
from django.urls import path
from django.shortcuts import render
from core.models import Video, Tag

class VideoAdmin(admin.ModelAdmin):
  def get_urls(self):
    urls = super().get_urls()
    custom_urls = [
      path('<int:id>/upload-video', self.upload_video, name='core_video_create')
    ]
    return custom_urls + urls

  def upload_video(self, request, id):
    return render(request, 'admin/core/upload_video.html')

admin.site.register(Video, VideoAdmin)
admin.site.register(Tag)