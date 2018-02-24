#!/usr/bin/env octave

setosa_label     = -1;
versicolor_label = +1;

training_set = csvread('data/training-set');

petals = training_set(:, 1);
sepals = training_set(:, 2);

setosa_petals = petals(training_set(:, 3) == setosa_label, :);
setosa_sepals = sepals(training_set(:, 3) == setosa_label, :);

versicolor_petals = petals(training_set(:, 3) == versicolor_label, :);
versicolor_sepals = sepals(training_set(:, 3) == versicolor_label, :);

hold on

h = plot(
    setosa_petals,     setosa_sepals,     'or',
    versicolor_petals, versicolor_sepals, 'xb'
);
legend([h], {'setosa', 'versicolor'});

hold off

axis ([
    min(petals) - 1,
    max(petals) + 1,
    min(sepals) - 1,
    max(sepals) + 1,
])

xlabel('Petal length, cm');
ylabel('Sepal length, cm');

title ('Training set');
legend('location', 'southeast');

saveas(1, 'img/training-set.png');
