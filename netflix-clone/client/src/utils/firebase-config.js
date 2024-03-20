import { initializeApp } from "firebase/app";
import {getAuth} from 'firebase/auth'

const firebaseConfig = {
  apiKey: "AIzaSyBSKrMCdSjwMw1txc8R243sIFHCnmuHIt4",
  authDomain: "netflix-project-3e03d.firebaseapp.com",
  projectId: "netflix-project-3e03d",
  storageBucket: "netflix-project-3e03d.appspot.com",
  messagingSenderId: "534786632144",
  appId: "1:534786632144:web:eb8e2a3da9955865a131d5",
  measurementId: "G-VR150CYMQ0"
};

const app = initializeApp(firebaseConfig);
export const firebaseAuth = getAuth(app)