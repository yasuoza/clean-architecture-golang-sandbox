DELETE FROM
  users
WHERE
  (
    name = "dummy-user-001"
    AND age = 12
  )
  OR (
    name = "dummy-user-002"
    AND age = 32
  )
  OR (
    name = "dummy-user-003"
    AND age = 2
  )
  OR (
    name = "dummy-user-004"
    AND age = 28
  )
  OR (
    name = "dummy-user-005"
    AND age = 54
  );
