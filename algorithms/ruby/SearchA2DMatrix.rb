def search_matrix(matrix, target)
  for row in matrix.reverse() do
      if target >= row[0] and target <= row.last
          for v in row do
              return true if v == target
          end
      end
  end

  return false
end