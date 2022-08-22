import {
  ChakraProvider,
  theme,
} from "@chakra-ui/react"
import { EditorState, convertToRaw } from 'draft-js';
import { Editor } from "react-draft-wysiwyg"
import "react-draft-wysiwyg/dist/react-draft-wysiwyg.css";

export const App = () => (
  <>
    <style>
        @import
        url(`https://fonts.googleapis.com/css2?family=Mulish&display=swap`);
      </style>``
    <Editor
      toolbar={{
        fontFamily: {
          options: ['Mulish']
        }
      }}
    />
  </>
)
