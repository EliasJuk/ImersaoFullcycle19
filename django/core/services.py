# Cores
YELLOW = "\033[33m"
RESET = "\033[0m"  # Reseta a cor

class VideoService: 

	def process_upload(self, video_id: int, chunk_index: int, chunk: bytes):
		print(f'{YELLOW}Processing upload for video {video_id}, chunk {chunk_index}...{RESET}')