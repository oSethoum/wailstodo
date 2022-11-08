import {
  Box,
  Button,
  Group,
  ScrollArea,
  TextInput,
  Text,
  Switch,
  Flex,
  Paper,
  Space,
  ActionIcon,
} from "@mantine/core";
import { useQuery, useQueryClient, useMutation } from "@tanstack/react-query";
import axios from "axios";
import { useEffect, useState } from "react";
import { IconTrash } from "@tabler/icons";

const client = new WebSocket("ws://localhost:5000/api/subscribe");

const fetchTodos = () =>
  axios.get("http://localhost:5000/api/todos").then((res) => res.data);

const createTodo = (data: { text: string }) =>
  axios.post("http://localhost:5000/api/todos", data);

const updateTodo = (todo: { id: number; text?: string; done?: boolean }) =>
  axios.patch(`http://localhost:5000/api/todos/${todo.id}`, todo);

const deleteTodo = (id: number) =>
  axios.delete(`http://localhost:5000/api/todos/${id}`);

function App() {
  const [text, setText] = useState("");
  const queryClient = useQueryClient();
  const { data, isSuccess } = useQuery({
    queryKey: ["todos"],
    queryFn: fetchTodos,
  });

  // Mutations
  const _create = useMutation({
    mutationFn: createTodo,
    onSuccess: () => {
      setText("");
    },
  });
  const _update = useMutation({ mutationFn: updateTodo });
  const _delete = useMutation({ mutationFn: deleteTodo });

  useEffect(() => {
    const onmessage = (e: MessageEvent<any>) => {
      queryClient.invalidateQueries(e.data);
    };
    client.addEventListener("message", onmessage);
    return () => {
      client.removeEventListener("message", onmessage);
    };
  }, []);

  return (
    <Flex
      sx={{ height: "100vh" }}
      justify="start"
      direction="column"
      align="center"
      p="lg"
    >
      <Group>
        <TextInput
          placeholder="todo"
          sx={{ width: 320 }}
          value={text}
          onChange={(e) => setText(e.currentTarget.value)}
        />
        <Button
          onClick={() => {
            _create.mutateAsync({
              text,
            });
          }}
          disabled={text.length < 5}
        >
          Add
        </Button>
      </Group>
      <Space h="md" />
      {isSuccess && (
        <ScrollArea sx={{ height: "400px" }}>
          {data.map((todo: any) => (
            <Paper
              key={todo.id}
              mb="xs"
              pb="sm"
              px="sm"
              withBorder
              sx={{ width: 400 }}
            >
              <Flex align="center" justify="space-between">
                <Switch
                  label={todo.text}
                  onChange={() => {
                    _update.mutateAsync({
                      id: todo.id,
                      done: !todo.done,
                    });
                  }}
                  checked={todo.done}
                />
                <ActionIcon
                  mt="sm"
                  color="red"
                  variant="subtle"
                  component={IconTrash}
                  onClick={(e) => {
                    _delete.mutateAsync(todo.id);
                  }}
                />
              </Flex>
            </Paper>
          ))}
        </ScrollArea>
      )}
    </Flex>
  );
}

export default App;
