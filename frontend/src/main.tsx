import React from "react";
import { createRoot } from "react-dom/client";
import "./style.css";
import App from "./App";
import { MantineProvider } from "@mantine/core";
import { QueryClientProvider, QueryClient } from "@tanstack/react-query";

const container = document.getElementById("root");
const queryClient = new QueryClient();
const root = createRoot(container!);

root.render(
  <React.StrictMode>
    <MantineProvider
      withGlobalStyles
      withNormalizeCSS
      theme={{ colorScheme: "dark" }}
    >
      <QueryClientProvider client={queryClient}>
        <App />
      </QueryClientProvider>
    </MantineProvider>
  </React.StrictMode>
);
