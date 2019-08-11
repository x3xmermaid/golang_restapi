use todo;

CREATE TABLE `category` (
  `id` int(30) NOT NULL,
  `name` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;



INSERT INTO `category` (`id`, `name`) VALUES
(1, 'PHP'),
(2, 'NODE js'),
(3, 'Golang'),
(4, 'Javascript');



CREATE TABLE `list` (
  `id` int(11) NOT NULL,
  `name` varchar(50) NOT NULL,
  `category` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;



INSERT INTO `list` (`id`, `name`, `category`) VALUES
(1, 'belajar', 2),
(2, 'Golang', 0),
(3, 'Golang', 1);


ALTER TABLE `category`
  ADD PRIMARY KEY (`id`);


ALTER TABLE `list`
  ADD PRIMARY KEY (`id`);


ALTER TABLE `category`
  MODIFY `id` int(30) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;


ALTER TABLE `list`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;