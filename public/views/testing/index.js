import EditorJS from './editor_tools/editor.js';
import Header from './editor_tools/header.js';
import List from './editor_tools/list.js';
import Embed from './editor_tools/embed.js';
import SimpleImage from './editor_tools/simple_image.js';
import Quote from './editor_tools/quote.js';

const editor = new EditorJS({
	holder: 'editorjs',

	tools: {
		header: {
			class: Header,
			inlineToolbar: ['link']
		},
		list: {
			class: List,
			inlineToolbar: [
				'link',
				'bold'
			]
		},
		embed: {
			class: Embed,
			inlineToolbar: false,
			config: {
				service: {
					youtube: true,
					coub: true,
					imgur: true,
					gfycat: true
				}
			}
		}, 
		simple_image: {
			class: SimpleImage,
		},
		quote: {
			class: Quote,
			inlineToolbar: true,
			shortcut: 'CMD+SHIFT+O',
			config: {
				quotePlaceholder: 'Enter a quote',
				captionPlaceholder: 'Quote\'s author',
			},
		}
	}
})

let saveBtn = document.querySelector('button');

saveBtn.addEventListener('click', function() {
	editor.save().then((outputdata) => {
		console.log('Article data: ', outputdata);
	}).catch((error) => {
		console.log('Saving failed: ', error)
	});
})